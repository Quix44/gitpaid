import type { Metadata } from "next";
import { Inter } from "next/font/google";

import Nav from "./components/Nav";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Gitpaid",
  description: "Development, Crowdsourced",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <Nav />
        {children}</body>
    </html>
  );
}
