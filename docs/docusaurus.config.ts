import type * as Preset from "@docusaurus/preset-classic";
import type { Config } from "@docusaurus/types";
import { themes as prismThemes } from "prism-react-renderer";

const config: Config = {
  title: "Gouse",
  tagline:
    "A modern essential Golang utility package delivering consistency, modularity, performance, & extras presets",
  favicon: "img/favicon.ico",

  // Set the production url of your site here
  url: "https://gouse.vercel.app",
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: "/",
  baseUrlIssueBanner: true,

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
      hideOnScroll: true,
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
          position: "right",
          className: "header-github-link",
          "aria-label": "GitHub repository",
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
              label: "Introduction",
              to: "/docs/intro",
            },
            {
              label: "Receipts",
              to: "/receipts",
            },
            {
              label: "Examples",
              to: "/examples",
            },
          ],
        },
        {
          title: "Links",
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
              label: "License",
              href: "https://github.com/thuongtruong1009/gouse/blob/main/LICENSE",
            },
            {
              label: "Change Log",
              href: "https://github.com/thuontruong1009/gouse/blob/main/CHANGELOG.md",
            },
            {
              label: "Funding",
              href: "https://fund.thuongtruong.me",
            },
          ],
        },
      ],
      copyright: `Copyright © 2024 Gouse ❤️`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
    // make Tailwind CSS and Docusaurus dark mode coexist
    // colorMode: {
    //   disableSwitch: true,
    //   respectPrefersColorScheme: true,
    // },
  } satisfies Preset.ThemeConfig,

  plugins: [
    async function myPlugin(context, options) {
      return {
        name: "docusaurus-tailwindcss",
        configurePostCss(postcssOptions) {
          postcssOptions.plugins.push(require("tailwindcss"));
          postcssOptions.plugins.push(require("autoprefixer"));
          return postcssOptions;
        },
      };
    },
  ],
};

export default config;
