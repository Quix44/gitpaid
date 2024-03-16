import { Button } from "@/components/ui/button"
import Image from 'next/image'


function Hero() {
    return (
        <section className="grid grid-cols-12 bg-primary  items-center rounded-3xl">
            <div className="col-span-5">
                <Image
                    src="/Octocatlogo.png"
                    width={500}
                    height={500}
                    alt="Picture of the author"
                />
            </div>
            <div className="col-span-6">
                <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl mb-6">
                    Code for Cash and Community Impact
                </h1>
                <p className="leading-7 [&:not(:first-child)]:mt-6">
                    Dive into the world where your coding skills pay off in more ways than one. Connect your passion for open-source with the thrill of earning crypto, one pull request at a time
                </p>
                <Button variant={"secondary"} className="mt-6">Sign up</Button>


            </div>

        </section >

    )
}

export default Hero