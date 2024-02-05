import type * as Preset from "@docusaurus/preset-classic";
import type { Config } from "@docusaurus/types";
import { themes as prismThemes } from "prism-react-renderer";

const config: Config = {
  title: "Gouse",
  tagline:
    "A modern essential Golang utility package delivering consistency, modularity, performance, & extras presets",
  favicon: "img/favicon.ico",

  // Set the production url of your site here
  url: "https://your-docusaurus-site.example.com",
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: "/",

  organizationName: "thuongtruong1009",
  projectName: "gouse",

  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",

  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      {
        docs: {
          sidebarPath: "./sidebars.ts",
          editUrl:
            "https://github.com/thuongtruong1009/gouse/tree/main/docs/main",
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/thuongtruong1009/gouse/tree/main/docs/main",
        },
        theme: {
          customCss: "./src/css/custom.css",
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    // Replace with your project's social card
    image: "img/docusaurus-social-card.jpg",
    navbar: {
      title: "Gouse",
      logo: {
        alt: "Gouse Logo",
        src: "img/logo.svg",
      },
      items: [
        {
          type: "docSidebar",
          sidebarId: "tutorialSidebar",
          position: "left",
          label: "Tutorial",
        },
        { to: "/blog", label: "Blog", position: "left" },
        {
          href: "https://github.com/thuongtruong1009/gouse",
          label: "GitHub",
          position: "right",
        },
      ],
    },
    footer: {
      style: "dark",
      links: [
        {
          title: "References",
          items: [
            {
              label: "Tutorial",
              to: "/docs/intro",
            },
            {
              label: "Receipts",
              to: "/receipts",
            },
          ],
        },
        {
          title: "More",
          items: [
            {
              label: "GitHub",
              href: "https://github.com/thuongtruong1009/gouse",
            },
            {
              label: "Dev.to",
              href: "https://dev.to/thuongtruong1009",
            },
            {
              label: "LinkedIn",
              href: "https://www.linkedin.com/in/thuongtruong1009/",
            },
          ],
        },
        {
          title: "More",
          items: [
            {
              label: "Blog",
              href: "https://blog.thuongtruong.me",
            },
            {
              label: "Awards",
              href: "https://awards.thuongtruong.me",
            },
          ],
        },
      ],
      copyright: `Copyright © 2024 Gouse`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
