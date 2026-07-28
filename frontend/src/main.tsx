import { createRoot } from "react-dom/client";
import { useState } from "react";
import "./style.css";
import App from "./App";
import { FluentProvider, webDarkTheme, webLightTheme } from "@fluentui/react-components";

function Root() {
    const [isDarkTheme, setIsDarkTheme] = useState(true);

    return (
        <FluentProvider theme={isDarkTheme ? webDarkTheme : webLightTheme} style={{ minHeight: "100vh" }}>
            <App
                isDarkTheme={isDarkTheme}
                onToggleTheme={() => setIsDarkTheme((currentTheme) => !currentTheme)}
            />
        </FluentProvider>
    );
}

const container = document.getElementById("root");
createRoot(container!).render(<Root />);
