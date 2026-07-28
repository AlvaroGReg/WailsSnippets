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
        <div className="snippets-table">
            {snippets.map((snippet) => (
                <div key={snippet.id} className="snippet-item" onClick={() => updateExample(snippet)}>
                    <span className="snippet-title">{snippet.title}</span>
                    <span className="snippet-lang">{snippet.language}</span>
                    <div className="snippet-buttons">
                        <button onClick={() => onDelete(snippet.id)}>Delete</button>
                    </div>
                </div>
            ))}
        </div>
    );
}

export default SnippetsTable;
