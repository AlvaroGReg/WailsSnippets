import {
    Button,
    Dialog,
    DialogActions,
    DialogBody,
    DialogContent,
    DialogSurface,
    DialogTitle,
} from "@fluentui/react-components";

type ConfirmDialogProps = {
    open: boolean;
    title: string;
    message: string;
    confirmLabel?: string;
    cancelLabel?: string;
    onClose: (confirmed: boolean) => void;
};

function ConfirmDialog({
    open,
    title,
    message,
    confirmLabel = "Confirm",
    cancelLabel = "Cancel",
    onClose,
}: ConfirmDialogProps) {
    return (
        <Dialog open={open} onOpenChange={(_, data) => {
            if (!data.open) {
                onClose(false);
            }
        }}>
            <DialogSurface>
                <DialogBody>
                    <DialogTitle>{title}</DialogTitle>
                    <DialogContent>{message}</DialogContent>
                    <DialogActions>
                        <Button onClick={() => onClose(false)}>{cancelLabel}</Button>
                        <Button appearance="primary" onClick={() => onClose(true)}>{confirmLabel}</Button>
                    </DialogActions>
                </DialogBody>
            </DialogSurface>
        </Dialog>
    );
}

export default ConfirmDialog;
