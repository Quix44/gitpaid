"use client"

import { CopyIcon } from "@radix-ui/react-icons"

import { Button } from "@/components/ui/button"
import {
    Dialog,
    DialogClose,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"


interface ConnectRepositoryButtonProps {
    connected: boolean;
}

export function ConnectRepositoryButton({ connected }: ConnectRepositoryButtonProps) {
    const buttonText = connected ? "Connected" : "Connect";
    const endpointURL = "https://api.emitly.dev/v1/webhook?listenerId=fn_258e7692473366ae283dac5e9fe09d00&apikey=tsxKKZgpDx7CM3yQq3pUYvXH4yp3oCe7KxaZtmqi"

    const handleCopy = async () => {
        try {
            // Make sure `valueToCopy` is the variable containing the string you want to copy
            await navigator.clipboard.writeText(endpointURL);
        } catch (err) {
            console.error('Failed to copy:', err);
        }
    };


    return (
        <Dialog>
            <DialogTrigger asChild>
                <Button variant="ghost" className={`${connected ? "bg-green-900 hover:bg-green-900" : "bg-red-900 hover:bg-red-900 animate-pulse"} `}>{buttonText}</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-md">
                <DialogHeader>
                    <DialogTitle>Connect Repository</DialogTitle>
                    <DialogDescription>
                        Place in your repository Settings {'>'} Webhooks and select {"Send me everything"}.
                    </DialogDescription>
                </DialogHeader>
                <div className="flex items-center space-x-2">
                    <div className="grid flex-1 gap-2">
                        <Label htmlFor="link" className="sr-only">
                            Link
                        </Label>
                        <Input
                            id="link"
                            defaultValue={endpointURL}
                            readOnly
                        />
                    </div>
                    <Button type="submit" size="sm" className="px-3" onClick={handleCopy}>
                        <span className="sr-only">Copy</span>
                        <CopyIcon className="h-4 w-4" />
                    </Button>
                </div>
                <DialogFooter className="sm:justify-start">
                    <DialogClose asChild>
                        <Button type="button" variant="secondary">
                            Close
                        </Button>
                    </DialogClose>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}

export default ConnectRepositoryButton
