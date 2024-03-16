import Hero from "./components/Hero";
import Searchinfo from "./components/Searchinfo";
import Statrow from "./components/Statrow";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <Hero />
      <Statrow />
      <Searchinfo />
    </main>
  );
}
