---

date: 2025-02-21T14:51:24+0800
layout: post

---

https://medium.com/@vidiptvashist/building-a-vector-database-from-scratch-in-python-6bd683ba5171

关键还是如何做 embedding。以及如何做相似度计算？

后面看一下有哪些常用的embedding model 和算法。

还有，代码中都是简单的字符串，如果特别长，怎么切片呢？这个参数是不是也重要？

```python
from typing import Any

import numpy as np


class VectorStore:
    def __init__(self):
        self.vector_data: dict[str, np.ndarray] = {}  # A dictionary to store vectors
        self.vector_index: dict[str, dict] = {}  # An indexing structure for retrieval

    def add_vector(self, vector_id: str, vector: np.ndarray):
        """
        Add a vector to the store.

        Args:
            vector_id (str or int): A unique identifier for the vector.
            vector (numpy.ndarray): The vector data to be stored.
        """
        self.vector_data[vector_id] = vector
        self._update_index(vector_id, vector)

    def get_vector(self, vector_id):
        """
        Retrieve a vector from the store.

        Args:
            vector_id (str or int): The identifier of the vector to retrieve.

        Returns:
            numpy.ndarray: The vector data if found, or None if not found.
        """
        return self.vector_data.get(vector_id)

    def _update_index(self, vector_id, vector):
        """
        Update the index with the new vector.

        Args:
            vector_id (str or int): The identifier of the vector.
            vector (numpy.ndarray): The vector data.
        """
        # In this simple example, we use brute-force cosine similarity for indexing
        for existing_id, existing_vector in self.vector_data.items():
            similarity = np.dot(vector, existing_vector) / (
                np.linalg.norm(vector) * np.linalg.norm(existing_vector)
            )
            if existing_id not in self.vector_index:
                self.vector_index[existing_id] = {}
            self.vector_index[existing_id][vector_id] = similarity

    def find_similar_vectors(self, query_vector, num_results=5):
        """
        Find similar vectors to the query vector using brute-force search.

        Args:
            query_vector (numpy.ndarray): The query vector for similarity search.
            num_results (int): The number of similar vectors to return.

        Returns:
            list: A list of (vector_id, similarity_score) tuples for the most similar vectors.
        """
        results: list[tuple[str, float]] = []
        for vector_id, vector in self.vector_data.items():
            similarity = np.dot(query_vector, vector) / (
                np.linalg.norm(query_vector) * np.linalg.norm(vector)
            )
            results.append((vector_id, similarity))

        # Sort by similarity in descending order
        results.sort(key=lambda x: x[1], reverse=True)

        # Return the top N results
        return results[:num_results]


# Establish a VectorStore instance
vector_store = VectorStore()  # Creating an instance of the VectorStore class

# Define sentences
sentences = [  # Defining a list of example sentences
    "I eat mango",
    "mango is my favorite fruit",
    "mango, apple, oranges are fruits",
    "fruits are good for health",
]

# Tokenization and Vocabulary Creation
vocabulary: set[str] = set()  # Initializing an empty set to store unique words
for sentence in sentences:  # Iterating over each sentence in the list
    tokens = (
        sentence.lower().split()
    )  # Tokenizing the sentence by splitting on whitespace and converting to lowercase
    vocabulary.update(tokens)  # Updating the set of vocabulary with unique tokens

# Assign unique indices to vocabulary words
word_to_index = {
    word: i for i, word in enumerate(vocabulary)
}  # Creating a dictionary mapping words to unique indices

# Vectorization

# Initializing an empty dictionary to store sentence vectors
sentence_vectors: dict[str, np.ndarray] = {}
for sentence in sentences:  # Iterating over each sentence in the list
    tokens = (
        sentence.lower().split()
    )  # Tokenizing the sentence by splitting on whitespace and converting to lowercase
    vector = np.zeros(
        len(vocabulary)
    )  # Initializing a numpy array of zeros for the sentence vector
    for token in tokens:  # Iterating over each token in the sentence
        vector[
            word_to_index[token]
        ] += 1  # Incrementing the count of the token in the vector
    sentence_vectors[sentence] = (
        vector  # Storing the vector for the sentence in the dictionary
    )

# Store in VectorStore
for sentence, vector in sentence_vectors.items():  # Iterating over each sentence vector
    vector_store.add_vector(
        sentence, vector
    )  # Adding the sentence vector to the VectorStore

# Similarity Search
query_sentence = "Mango is the best fruit"  # Defining a query sentence
query_vector = np.zeros(
    len(vocabulary)
)  # Initializing a numpy array of zeros for the query vector
query_tokens = (
    query_sentence.lower().split()
)  # Tokenizing the query sentence and converting to lowercase
for token in query_tokens:  # Iterating over each token in the query sentence
    if token in word_to_index:  # Checking if the token is present in the vocabulary
        query_vector[
            word_to_index[token]
        ] += 1  # Incrementing the count of the token in the query vector

similar_sentences = vector_store.find_similar_vectors(
    query_vector, num_results=2
)  # Finding similar sentences

# Display similar sentences
print("Query Sentence:", query_sentence)  # Printing the query sentence
print("Similar Sentences:")  # Printing the header for similar sentences
for (
    sentence,
    similarity,
) in similar_sentences:  # Iterating over each similar sentence and its similarity score
    print(
        f"{sentence}: Similarity = {similarity:.4f}"
    )  # Printing the similar sentence and its similarity score
```
