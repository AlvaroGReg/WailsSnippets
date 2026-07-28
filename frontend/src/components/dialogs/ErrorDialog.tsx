import {
    Button,
    Dialog,
    DialogActions,
    DialogBody,
    DialogContent,
    DialogSurface,
    DialogTitle,
} from "@fluentui/react-components";

type ErrorDialogProps = {
    error: string;
    onClose: () => void;
};

function ErrorDialog({ error, onClose }: ErrorDialogProps) {
    return (
        <Dialog
            open={Boolean(error)}
            modalType="alert"
            onOpenChange={(_, data) => {
                if (!data.open) {
                    onClose();
                }
            }}
        >
            <DialogSurface>
                <DialogBody>
                    <DialogTitle>Error</DialogTitle>
                    <DialogContent>{error}</DialogContent>
                    <DialogActions>
                        <Button appearance="primary" onClick={onClose}>Close</Button>
                    </DialogActions>
                </DialogBody>
            </DialogSurface>
        </Dialog>
    );
}

export default ErrorDialog;
