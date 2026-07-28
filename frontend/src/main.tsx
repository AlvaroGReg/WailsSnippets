import {createRoot} from 'react-dom/client'
import './style.css'
import App from './App'
import { FluentProvider, webDarkTheme } from "@fluentui/react-components";

const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
    <FluentProvider theme={webDarkTheme}>
        <App />
    </FluentProvider>,
)
