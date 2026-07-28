import { Button, Input, makeStyles, tokens } from "@fluentui/react-components";
import { DismissRegular, SearchRegular } from "@fluentui/react-icons";

type SearchBarProps = {
    value: string;
    onChange: (value: string) => void;
};

const useStyles = makeStyles({
    root: {
        width: "min(100%, 36rem)",
    },
    input: {
        backgroundColor: tokens.colorNeutralBackground1,
        borderRadius: tokens.borderRadiusMedium,
    },
});

function SearchBar({ value, onChange }: SearchBarProps) {
    const styles = useStyles();

    return (
        <div className={styles.root}>
            <Input
                aria-label="Search snippets"
                className={styles.input}
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
