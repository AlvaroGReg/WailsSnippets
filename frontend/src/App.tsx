import SnippetsTable from "./components/SnippetsTable";
import SearchBar from "./components/SearchBar";
import ConfirmDialog from "./components/dialogs/ConfirmDialog";
import ErrorDialog from "./components/dialogs/ErrorDialog";
import { Button } from "@fluentui/react-components";
import { useSnippets } from "./hooks/use-snippets";
import { useState } from "react";
import { AddRegular, BrightnessHighRegular, DarkThemeRegular } from "@fluentui/react-icons";

type AppProps = {
    isDarkTheme: boolean;
    onToggleTheme: () => void;
};

function App({ isDarkTheme, onToggleTheme }: AppProps) {
    const [searchQuery, setSearchQuery] = useState("");
    const [snippetPendingDeletion, setSnippetPendingDeletion] = useState<string | null>(null);
    const {
        snippets,
        error,
        clearError,
        storagePath,
        selectStorageDirectory,
        createSnippet,
        updateSnippet,
        deleteSnippet,
    } = useSnippets();

    function createExample() {
        const newSnippet = {
            title: "Sample title",
            language: "TypeScript",
            code: "console.log('Hello');",
            tags: ["typescript", "sample"],
        };

        void createSnippet(newSnippet);
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
                    onClick={createExample}
                    title="Create snippet"
                />
            </header>
            {/* <SnippetsTable
                snippets={snippets}
                onUpdate={updateSnippet}
                onDelete={deleteSnippet}
            /> */}
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
