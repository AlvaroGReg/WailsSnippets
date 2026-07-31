import { describe, expect, it, vi } from "vitest";
import userEvent from "@testing-library/user-event";
import { screen } from "@testing-library/react";
import SettingsDialog from "./SettingsDialog";
import { renderWithFluent } from "../../test/render";

describe("SettingsDialog", () => {
    it("changes the close-to-tray preference", async () => {
        const user = userEvent.setup();
        const onCloseToTrayChange = vi.fn();
        const onTraySnippetLimitChange = vi.fn();
        renderWithFluent(
            <SettingsDialog
                open
                closeToTrayEnabled={false}
                traySnippetLimit={5}
                onClose={vi.fn()}
                onCloseToTrayChange={onCloseToTrayChange}
                onTraySnippetLimitChange={onTraySnippetLimitChange}
            />,
        );

        await user.click(screen.getByRole("switch", { name: "Close to tray" }));

        expect(onCloseToTrayChange).toHaveBeenCalledWith(true);
    });

    it("changes the tray snippet limit", async () => {
        const user = userEvent.setup();
        const onTraySnippetLimitChange = vi.fn();
        renderWithFluent(
            <SettingsDialog
                open
                closeToTrayEnabled={false}
                traySnippetLimit={5}
                onClose={vi.fn()}
                onCloseToTrayChange={vi.fn()}
                onTraySnippetLimitChange={onTraySnippetLimitChange}
            />,
        );

        await user.clear(screen.getByRole("spinbutton", { name: "Snippets shown in tray" }));
        await user.type(screen.getByRole("spinbutton", { name: "Snippets shown in tray" }), "3");

        expect(onTraySnippetLimitChange).toHaveBeenCalledWith(3);
    });
});
