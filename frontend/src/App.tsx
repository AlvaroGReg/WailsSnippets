import SnippetsList from "./components/SnippetsList";
import SearchBar from "./components/SearchBar";
import ConfirmDialog from "./components/dialogs/ConfirmDialog";
import ErrorDialog from "./components/dialogs/ErrorDialog";
import { Button, Spinner } from "@fluentui/react-components";
import { useSnippets } from "./hooks/use-snippets";
import { useMemo, useState } from "react";
import { AddRegular, BrightnessHighRegular, DarkThemeRegular } from "@fluentui/react-icons";
import type { CreateSnippetInput, SnippetModel } from "./models/Snippet";

type AppProps = {
    isDarkTheme: boolean;
    onToggleTheme: () => void;
};

function App({ isDarkTheme, onToggleTheme }: AppProps) {
    const [searchQuery, setSearchQuery] = useState("");
    const [snippetPendingDeletion, setSnippetPendingDeletion] = useState<string | null>(null);
    const [snippetBeingEdited, setSnippetBeingEdited] = useState<SnippetModel | null | undefined>(undefined);
    const {
        snippets,
        error,
        clearError,
        isLoading,
        storagePath,
        selectStorageDirectory,
        createSnippet,
        updateSnippet,
        deleteSnippet,
    } = useSnippets();

    const filteredSnippets = useMemo(() => {
        const query = searchQuery.trim().toLocaleLowerCase();

        if (!query) {
            return snippets;
        }

        return snippets.filter((snippet) =>
            [snippet.title, snippet.language, snippet.code, ...snippet.tags]
                .some((field) => field.toLocaleLowerCase().includes(query)),
        );
    }, [searchQuery, snippets]);

    function saveSnippet(input: CreateSnippetInput) {
        if (snippetBeingEdited) {
            void updateSnippet({ ...snippetBeingEdited, ...input });
        } else {
            void createSnippet(input);
        }

        setSnippetBeingEdited(undefined);
    }

    function handleDeleteConfirmation(confirmed: boolean) {
        if (confirmed && snippetPendingDeletion) {
            void deleteSnippet(snippetPendingDeletion);
        }

        setSnippetPendingDeletion(null);
    }

    return (
        <main>
            <header>
                <SearchBar value={searchQuery} onChange={setSearchQuery} />
                <Button
                    appearance="primary"
                    aria-label="Create snippet"
                    icon={<AddRegular />}
                    onClick={() => setSnippetBeingEdited(null)}
                    title="Create snippet"
                />
            </header>
            {isLoading && <Spinner label="Loading snippets" />}
            {(!isLoading || snippets.length > 0) && (
                <SnippetsList
                    snippets={filteredSnippets}
                    onEdit={setSnippetBeingEdited}
                    onDelete={setSnippetPendingDeletion}
                />
            )}
            <footer>
                <Button
                    appearance="subtle"
                    className="storage-directory-button"
                    onClick={() => void selectStorageDirectory()}
                    title={storagePath || "No folder selected."}
                >
                    {storagePath || "No folder selected."}
                </Button>
                <Button
                    appearance="subtle"
                    className="theme-toggle-button"
                    icon={isDarkTheme ? <BrightnessHighRegular /> : <DarkThemeRegular />}
                    onClick={onToggleTheme}
                    aria-label={`Switch to ${isDarkTheme ? "light" : "dark"} theme`}
                    title={`Switch to ${isDarkTheme ? "light" : "dark"} theme`}
                />
            </footer>
            <ConfirmDialog
                open={snippetPendingDeletion !== null}
                title="Delete snippet"
                message="Are you sure you want to delete this snippet?"
                confirmLabel="Delete"
                onClose={handleDeleteConfirmation}
            />
            <ErrorDialog error={error} onClose={clearError} />
        </main>
    );
}

export default App;
