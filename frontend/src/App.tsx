import SnippetsTable from "./components/SnippetsTable";
import SearchBar from "./components/SearchBar";
import { useSnippets } from "./hooks/use-snippets";

function App() {
    const {
        snippets,
        error,
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

    return (
        <main>
            <head>
                <SearchBar></SearchBar>
                <button onClick={createExample}>+</button>
            </head>
            {error && <p>{error}</p>}
            <SnippetsTable
                snippets={snippets}
                onUpdate={updateSnippet}
                onDelete={deleteSnippet}
            />
            <footer>
                <button onClick={() => void selectStorageDirectory()}>{storagePath || "No folder selected."}</button>
                <button>Theme</button>
            </footer>
        </main>
    );
}

export default App;
