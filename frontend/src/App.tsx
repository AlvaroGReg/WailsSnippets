import { useEffect, useState } from "react";
import { CreateSnippet, DeleteSnippet, GetSnippets } from "../wailsjs/go/main/App";

function App() {
    const [snippets, setSnippets] = useState<any[]>([]);
    const [error, setError] = useState("");

    async function loadSnippets() {
        const result: any[] = await GetSnippets();
        console.log('loadsnippets::', result);
        setSnippets(result);
    }

    useEffect(() => {
        loadSnippets();
    }, []);

    async function createExample() {
        const newSnippet = {
            title: "Titulo",
            language: "TypeScript",
            code: "console.log('Hola');",
            tags: ["typescript", "ejemplo"],
        }

        console.log('createExample::', newSnippet);

        try {
            setError("");
            await CreateSnippet(newSnippet);
            await loadSnippets();
        } catch (err) {
            setError(err instanceof Error ? err.message : "No se pudo crear el snippet");
        }
    }

    async function deleteSnippet(id: string) {
        try {
            setError("");
            const result: any = await DeleteSnippet(id);
            await loadSnippets();
        } catch (err) {
            setError(err instanceof Error ? err.message : "No se pudo borrar el snippet");
        }
    }

    return (
        <main>
            <button onClick={createExample}>Crear snippet de ejemplo</button>

            {error && <p>{error}</p>}

            <ul>
                {snippets.map((snippet) => (
                    <li key={snippet.id}>
                        <span>{snippet.title}</span>
                        <span>{snippet.language}</span>
                        <button onClick={ () => deleteSnippet(snippet.id)}>
                            <span>Borrar</span>
                        </button>
                    </li>
                ))}
            </ul>
        </main>
    );
}

export default App;