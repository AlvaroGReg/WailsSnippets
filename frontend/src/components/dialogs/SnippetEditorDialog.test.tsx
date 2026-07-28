import { describe, expect, it, vi } from "vitest";
import userEvent from "@testing-library/user-event";
import { screen } from "@testing-library/react";
import SnippetEditorDialog from "./SnippetEditorDialog";
import { renderWithFluent } from "../../test/render";

describe("SnippetEditorDialog", () => {
    it("submits normalized tags for a new snippet", async () => {
        const user = userEvent.setup();
        const onSave = vi.fn();
        renderWithFluent(<SnippetEditorDialog open onClose={vi.fn()} onSave={onSave} />);

        await user.type(screen.getByRole("textbox", { name: "Title" }), "A snippet");
        await user.type(screen.getByRole("textbox", { name: "Language" }), "Go");
        await user.type(screen.getByRole("textbox", { name: "Code" }), "fmt.Println()");
        await user.type(screen.getByRole("textbox", { name: "Tags" }), " cli, go, ,testing ");
        await user.click(screen.getByRole("button", { name: "Create" }));

        expect(onSave).toHaveBeenCalledWith({
            title: "A snippet",
            language: "Go",
            code: "fmt.Println()",
            tags: ["cli", "go", "testing"],
        });
    });
});
