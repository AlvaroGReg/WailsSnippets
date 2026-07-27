import type { SnippetModel } from "../models/Snippet";

type SnippetsTableProps = {
    snippets: SnippetModel[];
    onUpdate: (snippet: SnippetModel) => void;
    onDelete: (id: string) => void;
};

function SnippetsTable({ snippets, onUpdate, onDelete }: SnippetsTableProps) {
    function updateExample(snippet: SnippetModel) {
        const updatedSnippet = {
            ...snippet,
            title: "Updated sample title",
            language: "TypeScript",
            code: "console.log('Updated sample');",
            tags: ["typescript", "updated-sample"],
        };

        void onUpdate(updatedSnippet);
    }

    return (
        <ul>
            {snippets.map((snippet) => (
                <li key={snippet.id}>
                    <span>{snippet.title}</span>
                    <span>{snippet.language}</span>
                    <button onClick={() => updateExample(snippet)}>
                        <span>Update</span>
                    </button>
                    <button onClick={() => onDelete(snippet.id)}>
                        <span>Delete</span>
                    </button>
                </li>
            ))}
        </ul>
    );
}

export default SnippetsTable;
