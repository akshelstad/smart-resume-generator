from openai import OpenAI
from dotenv import dotenv_values

config = dotenv_values(".env")
api_key = config["OPENAI_API_KEY"]

client = OpenAI(api_key=api_key)


def generate_resume(name, experience):
    prompt = f"Create a professional resume summary for {name} with experience in {experience}."

    response = client.responses.create(
        input=prompt,
        model="gpt-4",
    )

    return response.output[0].content[0].text