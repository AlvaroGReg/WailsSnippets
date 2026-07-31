import {
    Button,
    Dialog,
    DialogActions,
    DialogBody,
    DialogContent,
    DialogSurface,
    DialogTitle,
} from "@fluentui/react-components";

type StorageFileDialogProps = {
    open: boolean;
    onClose: () => void;
    onPickExisting: () => void;
    onCreateNew: () => void;
};

function StorageFileDialog({ open, onClose, onPickExisting, onCreateNew }: StorageFileDialogProps) {
    return (
        <Dialog open={open} onOpenChange={(_, data) => !data.open && onClose()}>
            <DialogSurface>
                <DialogBody>
                    <DialogTitle>Snippets file</DialogTitle>
                    <DialogContent>
                        Choose an existing JSON file or create a new one with the name you want.
                    </DialogContent>
                    <DialogActions>
                        <Button appearance="secondary" onClick={onClose}>Cancel</Button>
                        <Button appearance="secondary" onClick={onPickExisting}>Choose file</Button>
                        <Button appearance="primary" onClick={onCreateNew}>Create file</Button>
                    </DialogActions>
                </DialogBody>
            </DialogSurface>
        </Dialog>
    );
}

export default StorageFileDialog;
