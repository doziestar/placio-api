import os
from typing import Any, List, Optional, Dict
from fastapi import APIRouter, Depends, HTTPException, status, Header
from elevenlabs import generate, play, voices, clone, Voice

voice_router = APIRouter(prefix="/voice", tags=["voice"])

API_KEY = os.getenv("ELEVEN_API_KEY")


async def get_api_key(api_key: Optional[str] = Header(default=None)):
    if api_key != API_KEY:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Invalid API Key",
        )
    return api_key


@voice_router.get("/get_voices", response_model=List[Voice], dependencies=[Depends(get_api_key)])
async def get_voices() -> Any:
    try:
        available_voices = voices().voices
        return available_voices
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@voice_router.get("/get_voice", response_model=Voice, dependencies=[Depends(get_api_key)])
async def get_voice(voice_id: str) -> Any:
    try:
        available_voices = voices().voices
        for voice in available_voices:
            if voice.voiceId == voice_id:
                return voice
        raise HTTPException(status_code=404, detail="Voice not found")
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@voice_router.post("/generate", response_model=Voice, dependencies=[Depends(get_api_key)])
async def generate_voice(text: str, voice: str) -> Any:
    try:
        generated_voice = generate(text, voice)
        return generated_voice
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@voice_router.post("/play", response_model=str, dependencies=[Depends(get_api_key)])
async def play_voice(text: bytes, voice: str) -> Any:
    try:
        return play(text)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@voice_router.post("/generate_and_play", response_model=str, dependencies=[Depends(get_api_key)])
async def generate_and_play_voice(text: str, voice: str) -> Any:
    try:
        generated_text = generate(text, voice)
        return play(generated_text)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@voice_router.post("/clone_voice", response_model=str, dependencies=[Depends(get_api_key)])
async def clone_voice(text: str, name: str, files: List[str]) -> Any:
    try:
        cloned_voice = clone(name=name, text=text, files=files)
        return cloned_voice
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@voice_router.post("/generate_and_clone", response_model=str, dependencies=[Depends(get_api_key)])
async def generate_and_clone_voice(text: str, name: str, files: List[str], voice: str) -> Any:
    try:
        generated_text = generate(text, voice)
        cloned_voice = clone(name=name, text=generated_text, files=files)
        return cloned_voice
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@voice_router.post("/generate_and_play_and_clone", response_model=str, dependencies=[Depends(get_api_key)])
async def generate_and_play_and_clone_voice(text: str, name: str, files: List[str], voice: str) -> Any:
    try:
        generated_text = generate(text, voice)
        cloned_voice = clone(name=name, text=play(generated_text), files=files)
        return cloned_voice
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
