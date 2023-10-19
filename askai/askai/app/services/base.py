import os

import openai
from typing import Dict, Optional

API_KEY = os.getenv("OPENAI_API_KEY")

# Initialize the openai package with your API key
openai.api_key = API_KEY


class OpenAIApi:
    def ask_openai(self, prompt: str, max_tokens: int = 1024, temperature: float = 0.6) -> Dict[str, Optional[str]]:
        try:
            # Start with a system message to guide the model
            messages = [
              {"role": "system", "content": "You are a helpful assistant. Provide accurate and concise answers."},
              {"role": "user", "content": prompt}
            ]

            completion = openai.ChatCompletion.create(
              model="gpt-3.5-turbo",
              messages=messages,
              temperature=temperature,
              max_tokens=max_tokens,
            )

            response_content = completion.choices[0].message.content

            # Optionally check for minimum length
            # if len(response_content.split()) < MIN_TOKENS:
            #     return {"success": False, "error": "Response too short"}

            return {
              "success": True,
              "data": response_content
            }

        except openai.OpenAIError as e:
            # Handle specific OpenAI errors here, like rate limits
            return {
              "success": False,
              "error": str(e)
            }
        except Exception as e:
            # General error handling
            return {
              "success": False,
              "error": f"An unexpected error occurred: {str(e)}"
            }


if __name__ == "__main__":
    api = OpenAIApi()
    response = api.ask_openai("What is photosynthesis?")
    if response["success"]:
        print(response["data"])
    else:
        print("Error:", response["error"])
