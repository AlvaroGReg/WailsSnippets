export type SnippetModel = {
    id: string
    title: string
    language: string
    code: string
    tags: string[]
    createdAt: string
}

export type CreateSnippetInput = Omit<SnippetModel, "id" | "createdAt">;
