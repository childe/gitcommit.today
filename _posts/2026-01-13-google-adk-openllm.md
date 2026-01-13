---

date: 2026-01-13T18:36:20+0800
layout: post

---

```python
import os
from google.adk.agents import Agent
from google.adk.models.lite_llm import LiteLlm  # For multi-model support
import litellm  # Import for proxy configuration

os.environ["LITELLM_PROXY_API_KEY"] = ""
os.environ["LITELLM_PROXY_API_BASE"] = ""

litellm.use_litellm_proxy = True

root_agent = Agent(
    model=LiteLlm(model="gemini-2.5-pro"),
    name="root_agent",
    description="You are a helpful assistant that tells the current time in cities.",
    instruction="You are a helpful assistant that tells the current time in cities.",
    # tools=[get_current_time],
)
```
