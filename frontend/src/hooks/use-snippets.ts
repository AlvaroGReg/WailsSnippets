import { useCallback, useEffect, useState } from "react";
import type { CreateSnippetInput, SnippetModel } from "../models/Snippet";
import * as snippetsService from "../services/snippets-service";

function getErrorMessage(error: unknown): string {
    return error instanceof Error ? error.message : "An unexpected error occurred.";
}

export function useSnippets() {
    const [snippets, setSnippets] = useState<SnippetModel[]>([]);
    const [error, setError] = useState("");
    const [storagePath, setStoragePath] = useState("");

    const loadSnippets = useCallback(async () => {
        try {
            setError("");
            setSnippets(await snippetsService.getSnippets());
        } catch (error) {
            setError(getErrorMessage(error));
        }
    }, []);

    useEffect(() => {
        void loadSnippets();
        void snippetsService.getSnippetsStoragePath().then(setStoragePath).catch((error: unknown) => {
            setError(getErrorMessage(error));
        });
    }, [loadSnippets]);

    const selectStorageDirectory = useCallback(async () => {
        try {
            setError("");
            setStoragePath(await snippetsService.selectSnippetsDirectory());
            await loadSnippets();
        } catch (error) {
            setError(getErrorMessage(error));
        }
    }, [loadSnippets]);

    const createSnippet = useCallback(async (input: CreateSnippetInput) => {
        try {
            setError("");
            const snippet = await snippetsService.createSnippet(input);
            setSnippets((currentSnippets) => [...currentSnippets, snippet]);
        } catch (error) {
            setError(getErrorMessage(error));
        }
    }, []);

    const updateSnippet = useCallback(async (snippet: SnippetModel) => {
        try {
            setError("");
            const updatedSnippet = await snippetsService.updateSnippet(snippet);
            setSnippets((currentSnippets) =>
                currentSnippets.map((currentSnippet) =>
                    currentSnippet.id === updatedSnippet.id ? updatedSnippet : currentSnippet,
                ),
            );
        } catch (error) {
            setError(getErrorMessage(error));
        }
    }, []);

    const deleteSnippet = useCallback(async (id: string) => {
        try {
            setError("");
            await snippetsService.deleteSnippet(id);
            setSnippets((currentSnippets) =>
                currentSnippets.filter((snippet) => snippet.id !== id),
            );
        } catch (error) {
            setError(getErrorMessage(error));
        }
    }, []);

    return {
        snippets,
        error,
        storagePath,
        selectStorageDirectory,
        createSnippet,
        updateSnippet,
        deleteSnippet,
    };
}
