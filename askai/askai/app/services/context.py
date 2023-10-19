
# For the /ask endpoint
def ask_question_context(question: str) -> str:
    return f"Provide a comprehensive and concise answer to the following question: '{question}'."


# For the /explore/{topic} endpoint
def explore_topic_context(topic: str) -> str:
    return f"Give a brief summary about the topic '{topic}' and suggest some closely related topics."


# For the /explore/{topic}/resources endpoint
def explore_topic_resources_context(topic: str) -> str:
    return f"List some reputable resources (books, websites, papers) related to the topic '{topic}'."


# For the /practice_language endpoint
def practice_language_context(text: str, language: str) -> str:
    return f"Review the following text in {language} and provide corrections or suggestions for improvement: '{text}'."


# For the /homework_helper endpoint
def homework_helper_context(question: str, hint_only: bool) -> str:
    if hint_only:
        return f"Provide a hint or guidance for solving the following problem: '{question}'."
    else:
        return f"Solve the following problem and provide a clear explanation: '{question}'."


# For the /essay_grader endpoint
def essay_grader_context(essay: str) -> str:
    return f"Review the following essay and provide a grade (A+ to F) based on its content, coherence, and grammar:\n\n{essay}"


# For the /generate_flashcards endpoint
def generate_flashcards_context(topic: str) -> str:
    return f"Generate flashcards for the topic '{topic}', including questions and answers covering key points."


# For the /simulate/{scenario} endpoint
def simulate_scenario_context(scenario: str) -> str:
    return f"Given the scenario '{scenario}', describe a potential outcome or response to simulate this situation."


# For the /game/{game_type} endpoint
def educational_game_context(game_type: str) -> str:
    return f"Initiate and manage a game of type '{game_type}'. Provide instructions, rules, and potential challenges."


# For the /discussion_room/{room_id}/message endpoint
def post_message_context(room_id: int, message: str) -> str:
    return f"In discussion room {room_id}, a user wants to convey the following message: '{message}'. Analyze its content and intent."


# For the /learning_path endpoint
def create_learning_path_context(subject: str, objectives: list) -> str:
    objectives_str = ', '.join(objectives)
    return f"Design a learning path for the subject '{subject}' with the following objectives: {objectives_str}. Provide a structured outline and potential resources."


# For the /progress/{user_id} endpoint
def track_progress_context(user_id: int) -> str:
    return f"Retrieve progress details for user with ID {user_id}. Analyze and present data related to their learning milestones, achievements, and areas of improvement."
