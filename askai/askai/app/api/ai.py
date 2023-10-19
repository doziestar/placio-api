from typing import Any, List

from fastapi import APIRouter

from app.services.base import OpenAIApi
from app.services.context import ask_question_context, explore_topic_context, explore_topic_resources_context, \
  practice_language_context, homework_helper_context, essay_grader_context, generate_flashcards_context, \
  simulate_scenario_context, educational_game_context, post_message_context, create_learning_path_context, \
  track_progress_context

ai_router = APIRouter(prefix="/ai", tags=["ai"])

openai_api = OpenAIApi()


@ai_router.get("/hello")
def hello() -> Any:
    return {"message": "Hello, World!"}


@ai_router.post("/ask")
def ask_question(question: str) -> Any:
    context = ask_question_context(question)
    response = openai_api.ask_openai(context)
    if response["success"]:
        return {"answer": response["data"]}
    else:
        return {"error": response["error"]}


@ai_router.get("/explore/{topic}")
def explore_topic(topic: str) -> Any:
    context = explore_topic_context(topic)
    response = openai_api.ask_openai(context)
    if response["success"]:
        # Assuming the model returns in format: "Summary: ... Related topics: ..."
        summary, related_topics = response["data"].split("Related topics:")
        return {"summary": summary.strip(), "related_topics": [topic.strip() for topic in related_topics.split(",")]}
    else:
        return {"error": response["error"]}


@ai_router.get("/explore/{topic}/resources")
def explore_topic_resources(topic: str) -> Any:
    context = explore_topic_resources_context(topic)
    response = openai_api.ask_openai(context)
    if response["success"]:
        # Example output parsing: "Resources: resource1, resource2, ..."
        resources = response["data"].split("Resources:")[1].split(",")
        return {"resources": [res.strip() for res in resources]}
    else:
        return {"error": response["error"]}


@ai_router.post("/practice_language")
def practice_language(text: str, language: str) -> Any:
    context = practice_language_context(text, language)
    response = openai_api.ask_openai(context)
    if response["success"]:
        # The corrections/suggestions might be in the response
        return {"corrections": response["data"]}
    else:
        return {"error": response["error"]}


@ai_router.post("/homework_helper")
def homework_helper(question: str, hint_only: bool = False) -> Any:
    context = homework_helper_context(question, hint_only)
    response = openai_api.ask_openai(context)
    if response["success"]:
        return {"response": response["data"]}
    else:
        return {"error": response["error"]}


@ai_router.post("/essay_grader")
def essay_grader(essay: str) -> Any:
    context = essay_grader_context(essay)
    response = openai_api.ask_openai(context)
    if response["success"]:
        # The grade might be at the start or end of the response
        return {"grade": response["data"].split()[0]}
    else:
        return {"error": response["error"]}


@ai_router.post("/generate_flashcards")
def generate_flashcards(topic: str) -> Any:
    context = generate_flashcards_context(topic)
    response = openai_api.ask_openai(context)
    if response["success"]:
        # Parsing for "Flashcards: flashcard1, flashcard2, ..."
        flashcards = response["data"].split("Flashcards:")[1].split(",")
        return {"flashcards": [flashcard.strip() for flashcard in flashcards]}
    else:
        return {"error": response["error"]}


@ai_router.get("/simulate/{scenario}")
def simulate_scenario(scenario: str) -> Any:
    context = simulate_scenario_context(scenario)
    response = openai_api.ask_openai(context)
    if response["success"]:
        return {"simulation_result": response["data"]}
    else:
        return {"error": response["error"]}


@ai_router.get("/game/{game_type}")
def educational_game(game_type: str) -> Any:
    context = educational_game_context(game_type)
    response = openai_api.ask_openai(context)
    if response["success"]:
        return {"game_data": response["data"]}
    else:
        return {"error": response["error"]}


@ai_router.post("/discussion_room/{room_id}/message")
def post_message(room_id: int, message: str) -> Any:
    context = post_message_context(room_id, message)
    response = openai_api.ask_openai(context)
    if response["success"]:
        return {"status": response["data"]}
    else:
        return {"error": response["error"]}


@ai_router.post("/learning_path")
def create_learning_path(subject: str, objectives: List[str]) -> Any:
    context = create_learning_path_context(subject, objectives)
    response = openai_api.ask_openai(context)
    if response["success"]:
        return {"path_data": response["data"]}
    else:
        return {"error": response["error"]}


@ai_router.get("/progress/{user_id}")
def track_progress(user_id: int) -> Any:
    context = track_progress_context(user_id)
    response = openai_api.ask_openai(context)
    if response["success"]:
        return {"progress_data": response["data"]}
    else:
        return {"error": response["error"]}



