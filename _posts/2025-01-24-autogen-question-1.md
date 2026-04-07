---

date: 2025-01-24T11:59:21+0800
title: AutoGen SelectorGroupChat的participants参数含义
layout: post

---

autogen里面的一段代码如下：

```python

selector_group_chat = SelectorGroupChat(

[add_agent, multiply_agent, subtract_agent, divide_agent, identity_agent],

model_client=OpenAIChatCompletionClient(model="gpt-4o"),

termination_condition=termination_condition,

allow_repeated_speaker=True, # Allow the same agent to speak multiple times, necessary for this task.

selector_prompt=(

"Available roles:\n{roles}\nTheir job descriptions:\n{participants}\n"

"Current conversation history:\n{history}\n"

"Please select the most appropriate role for the next message, and only return the role name."

),

)

```

selector_prompt 里面的 participants 是什么意思？
