import {
    Button,
    Dialog,
    DialogActions,
    DialogBody,
    DialogContent,
    DialogSurface,
    DialogTitle,
    Field,
    Input,
    Switch,
} from "@fluentui/react-components";

type SettingsDialogProps = {
    open: boolean;
    closeToTrayEnabled: boolean;
    traySnippetLimit: number;
    onClose: () => void;
    onCloseToTrayChange: (enabled: boolean) => void;
    onTraySnippetLimitChange: (limit: number) => void;
};

function SettingsDialog({
    open,
    closeToTrayEnabled,
    traySnippetLimit,
    onClose,
    onCloseToTrayChange,
    onTraySnippetLimitChange,
}: SettingsDialogProps) {
    return (
        <Dialog open={open} onOpenChange={(_, data) => !data.open && onClose()}>
            <DialogSurface>
                <DialogBody>
                    <DialogTitle>Settings</DialogTitle>
                    <DialogContent>
                        <Switch
                            checked={closeToTrayEnabled}
                            label="Close to tray"
                            onChange={(_, data) => onCloseToTrayChange(data.checked)}
                        />
                        <p>Keep SnippetsDome running in the notification area when its window is closed.</p>
                        <Field label="Snippets shown in tray">
                            <Input
                                aria-label="Snippets shown in tray"
                                min={1}
                                type="number"
                                value={String(traySnippetLimit)}
                                onChange={(_, data) => {
                                    const limit = Number(data.value);
                                    if (Number.isInteger(limit) && limit > 0) {
                                        onTraySnippetLimitChange(limit);
                                    }
                                }}
                            />
                        </Field>
                    </DialogContent>
                    <DialogActions>
                        <Button appearance="primary" onClick={onClose}>Done</Button>
                    </DialogActions>
                </DialogBody>
            </DialogSurface>
        </Dialog>
    );
}

export default SettingsDialog;
