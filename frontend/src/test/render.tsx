import { FluentProvider, webLightTheme } from "@fluentui/react-components";
import { render } from "@testing-library/react";
import type { ReactElement } from "react";

export function renderWithFluent(ui: ReactElement) {
    return render(<FluentProvider theme={webLightTheme}>{ui}</FluentProvider>);
}
