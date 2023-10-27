
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
    language_dict = {
        'French': 'bonjour le monde',
        'Spanish': 'hola mundo',
        'German': 'hallo Welt',
        'Italian': 'ciao mondo',
        'Portuguese': 'olá mundo',
        'Russian': 'привет мир',
        'Chinese (Simplified)': '你好，世界',
        'Japanese': 'こんにちは世界',
        'Korean': '안녕하세요 세계',
        'Arabic': 'مرحبا بالعالم',
        'Hindi': 'नमस्ते दुनिया',
        'Bengali': 'হ্যালো বিশ্ব',
        'Turkish': 'merhaba dünya',
        'Vietnamese': 'xin chào thế giới',
        'Thai': 'สวัสดีโลก',
        'Indonesian': 'halo dunia',
        'Filipino': 'kamusta mundo',
    }
    translation_example = language_dict.get(language, 'Translation not available')
    return f"""
    You have been given a text in English: '{text}'.
    As a native {language} speaker, please translate the text '{text}' into {language}.
    Example translation: '{translation_example}'.

    Note: The text may not be grammatically correct or make sense in the given language.
    I am looking for a translation that is as close as possible to the original text.
    I am not looking for a perfect translation, but rather a translation that is natural and fluent.
    I will be evaluating your translation based on the following criteria:
    - Fluency: Does the translation sound natural and fluent?
    - Accuracy: Does the translation accurately convey the meaning of the original text?
    - Grammar: Does the translation use proper grammar and punctuation?
    - Spelling: Does the translation use proper spelling?
    - Vocabulary: Does the translation use proper vocabulary?
    - Style: Does the translation use proper style and tone?

    Important: Please do add any additional text to the translation.
    Please only translate the text provided.

    Please translate the text '{text}' into {language}.
    """

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
