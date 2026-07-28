import { Badge, Button, Toast, ToastTitle, Toaster, useToastController } from "@fluentui/react-components";
import type { SnippetModel } from "../models/Snippet";
import "./SnippetsList.css";

type SnippetsListProps = {
    snippets: SnippetModel[];
    onEdit: (snippet: SnippetModel) => void;
    onDelete: (id: string) => void;
};

function SnippetsList({ snippets, onEdit, onDelete }: SnippetsListProps) {
    const toasterId = "snippets-list";
    const { dispatchToast } = useToastController(toasterId);

    async function addToClipboard(code: string) {
        try {
            await navigator.clipboard.writeText(code);
            dispatchToast(
                <Toast>
                    <ToastTitle>Code copied to clipboard</ToastTitle>
                </Toast>,
                { intent: "success" },
            );
        } catch {
            dispatchToast(
                <Toast>
                    <ToastTitle>Could not copy the code</ToastTitle>
                </Toast>,
                { intent: "error" },
            );
        }
    }

    return (
        <div className="snippets-list">
            <Toaster toasterId={toasterId} position="bottom-end" />
            {snippets.length === 0 ? (
                <p className="snippets-list-empty">Empty list</p>
            ) : snippets.map((snippet) => (
                <article key={snippet.id} className="snippet-item">
                    <div className="snippet-head">
                        <span className="snippet-title">{snippet.title}</span>
                        <Button appearance="subtle" onClick={() => void addToClipboard(snippet.code)}>Copy</Button>
                    </div>
                    <div className="snippet-subtitle">
                        <span className="snippet-lang">{snippet.language}</span>
                        <div className="snippet-actions">
                            <Button appearance="outline" onClick={() => onDelete(snippet.id)}>Delete</Button>
                            <Button appearance="subtle" onClick={() => onEdit(snippet)}>Edit</Button>
                        </div>
                    </div>
                    <div className="snippet-body">
                        <pre className="snippet-code"><code>{snippet.code}</code></pre>
                    </div>
                    <ul className="snippet-tags">
                        {snippet.tags.map((tag) => (
                            <li key={`${snippet.id}-${tag}`}>
                                <Badge appearance="tint" color="brand">{tag}</Badge>
                            </li>
                        ))}
                    </ul>
                </article>
            ))}
        </div>
    );
}

export default SnippetsList;
