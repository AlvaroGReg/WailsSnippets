import { describe, expect, it, vi } from "vitest";
import userEvent from "@testing-library/user-event";
import { screen } from "@testing-library/react";
import { useState } from "react";
import SearchBar from "./SearchBar";
import { renderWithFluent } from "../../test/render";

function SearchBarHarness({ onChange }: { onChange: (value: string) => void }) {
    const [value, setValue] = useState("");

    function handleChange(nextValue: string) {
        setValue(nextValue);
        onChange(nextValue);
    }

    return <SearchBar value={value} onChange={handleChange} />;
}

describe("SearchBar", () => {
    it("notifies changes and clears a non-empty query", async () => {
        const user = userEvent.setup();
        const onChange = vi.fn();
        renderWithFluent(<SearchBarHarness onChange={onChange} />);

        await user.type(screen.getByRole("textbox", { name: "Search snippets" }), "react");
        expect(onChange).toHaveBeenLastCalledWith("react");

        await user.click(screen.getByRole("button", { name: "Clear search" }));
        expect(onChange).toHaveBeenLastCalledWith("");
    });
});
