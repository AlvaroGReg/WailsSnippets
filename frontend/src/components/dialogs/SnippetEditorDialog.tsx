import {
    Button,
    Dialog,
    DialogActions,
    DialogBody,
    DialogContent,
    DialogSurface,
    DialogTitle,
    Field,
    Input,
    Textarea,
} from "@fluentui/react-components";
import { useEffect, useState } from "react";
import type { CreateSnippetInput, SnippetModel } from "../../models/Snippet";

type SnippetEditorDialogProps = {
    open: boolean;
    snippet?: SnippetModel;
    onClose: () => void;
    onSave: (input: CreateSnippetInput) => void;
};

const emptySnippet: CreateSnippetInput = {
    title: "",
    language: "",
    code: "",
    tags: [],
};

function SnippetEditorDialog({ open, snippet, onClose, onSave }: SnippetEditorDialogProps) {
    const [form, setForm] = useState<CreateSnippetInput>(emptySnippet);
    const [tags, setTags] = useState("");

    useEffect(() => {
        if (!open) {
            return;
        }

        setForm(snippet ? {
            title: snippet.title,
            language: snippet.language,
            code: snippet.code,
            tags: snippet.tags,
        } : emptySnippet);
        setTags(snippet?.tags.join(", ") ?? "");
    }, [open, snippet]);

    function updateField(field: keyof Omit<CreateSnippetInput, "tags">, value: string) {
        setForm((current) => ({ ...current, [field]: value }));
    }

    function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        onSave({
            ...form,
            tags: tags.split(",").map((tag) => tag.trim()).filter(Boolean),
        });
    }

    const isEditing = Boolean(snippet);

    return (
        <Dialog open={open} onOpenChange={(_, data) => !data.open && onClose()}>
            <DialogSurface>
                <form onSubmit={handleSubmit}>
                    <DialogBody>
                        <DialogTitle>{isEditing ? "Edit snippet" : "Create snippet"}</DialogTitle>
                        <DialogContent>
                            <Field label="Title" required>
                                <Input
                                    value={form.title}
                                    onChange={(_, data) => updateField("title", data.value)}
                                    required
                                    autoFocus
                                />
                            </Field>
                            <Field label="Language" required>
                                <Input
                                    value={form.language}
                                    onChange={(_, data) => updateField("language", data.value)}
                                    required
                                />
                            </Field>
                            <Field label="Code" required>
                                <Textarea
                                    value={form.code}
                                    onChange={(_, data) => updateField("code", data.value)}
                                    required
                                    resize="vertical"
                                    rows={10}
                                />
                            </Field>
                            <Field label="Tags" hint="Separate tags with commas">
                                <Input value={tags} onChange={(_, data) => setTags(data.value)} />
                            </Field>
                        </DialogContent>
                        <DialogActions>
                            <Button onClick={onClose}>Cancel</Button>
                            <Button appearance="primary" type="submit">
                                {isEditing ? "Save" : "Create"}
                            </Button>
                        </DialogActions>
                    </DialogBody>
                </form>
            </DialogSurface>
        </Dialog>
    );
}

export default SnippetEditorDialog;
