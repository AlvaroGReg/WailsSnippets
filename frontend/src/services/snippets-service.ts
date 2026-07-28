import {
    CreateSnippet,
    DeleteSnippet,
    GetSnippets,
	GetSnippetsStoragePath,
	SelectSnippetsDirectory,
    UpdateSnippet,
} from "../../wailsjs/go/main/App";
import type { CreateSnippetInput, SnippetModel } from "../models/Snippet";

class SnippetsServiceError extends Error {
    constructor(message: string, cause: unknown) {
        super(message);
        this.name = "SnippetsServiceError";
        this.cause = cause;
    }
}

async function runRequest<T>(request: () => Promise<T>, errorMessage: string): Promise<T> {
    try {
        return await request();
    } catch (error) {
        throw new SnippetsServiceError(errorMessage, error);
    }
}

export function getSnippets(): Promise<SnippetModel[]> {
    return runRequest(GetSnippets, "No hay una ruta configurada. Selecciona una carpeta de snippets.");
}

export function createSnippet(input: CreateSnippetInput): Promise<SnippetModel> {
    return runRequest(() => CreateSnippet(input), "Unable to create the snippet.");
}

export function updateSnippet(snippet: SnippetModel): Promise<SnippetModel> {
    return runRequest(() => UpdateSnippet(snippet), "Unable to update the snippet.");
}

export function deleteSnippet(id: string): Promise<void> {
    return runRequest(() => DeleteSnippet(id), "Unable to delete the snippet.");
}

export function getSnippetsStoragePath(): Promise<string> {
    return runRequest(GetSnippetsStoragePath, "Unable to get the snippets file path.");
}

export function selectSnippetsDirectory(): Promise<string> {
    return runRequest(SelectSnippetsDirectory, "Unable to select the snippets folder.");
}
