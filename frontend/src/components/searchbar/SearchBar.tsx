import './SearchBar.css'
import { Button, Input } from "@fluentui/react-components";
import { DismissRegular, SearchRegular } from "@fluentui/react-icons";

type SearchBarProps = {
    value: string;
    onChange: (value: string) => void;
};

function SearchBar({ value, onChange }: SearchBarProps) {
    return (
        <div className="search-bar">
            <Input
                aria-label="Search snippets"
                className="search-input"
                contentBefore={<SearchRegular aria-hidden="true" />}
                contentAfter={value ? (
                    <Button
                        appearance="transparent"
                        aria-label="Clear search"
                        icon={<DismissRegular />}
                        onClick={() => onChange("")}
                        size="small"
                    />
                ) : undefined}
                onChange={(_, data) => onChange(data.value)}
                placeholder="Search snippets"
                value={value}
            />
        </div>
    );
}

export default SearchBar;
