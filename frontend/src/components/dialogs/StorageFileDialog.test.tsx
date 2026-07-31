import { describe, expect, it, vi } from "vitest";
import userEvent from "@testing-library/user-event";
import { screen } from "@testing-library/react";
import StorageFileDialog from "./StorageFileDialog";
import { renderWithFluent } from "../../test/render";

describe("StorageFileDialog", () => {
    it("offers existing-file and create-file actions", async () => {
        const user = userEvent.setup();
        const onClose = vi.fn();
        const onPickExisting = vi.fn();
        const onCreateNew = vi.fn();
        renderWithFluent(
            <StorageFileDialog
                open
                onClose={onClose}
                onPickExisting={onPickExisting}
                onCreateNew={onCreateNew}
            />,
        );

        await user.click(screen.getByRole("button", { name: "Choose file" }));
        expect(onPickExisting).toHaveBeenCalledOnce();

        await user.click(screen.getByRole("button", { name: "Create file" }));
        expect(onCreateNew).toHaveBeenCalledOnce();

        await user.click(screen.getByRole("button", { name: "Cancel" }));
        expect(onClose).toHaveBeenCalledOnce();
    });
});
