import SnippetsTable from "./components/snippets-table";
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
            <section className="storage-settings">
                <button onClick={() => void selectStorageDirectory()}>
                    Choose folder to save snippets
                </button>
                <p>{storagePath || "No folder selected."}</p>
            </section>
            <button onClick={createExample}>Create sample snippet</button>

            {error && <p>{error}</p>}

            <SnippetsTable
                snippets={snippets}
                onUpdate={updateSnippet}
                onDelete={deleteSnippet}
            />
        </main>
    );
}

export default App;
