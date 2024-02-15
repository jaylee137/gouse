import Link from "@docusaurus/Link";
import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import HomepageFeatures from "@site/src/components/HomepageFeatures";
import Heading from "@theme/Heading";
import Layout from "@theme/Layout";
import clsx from "clsx";

import styles from "./index.module.css";

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  const buttons = [
    {
      label: "Get started",
      link: "/docs/intro",
    },
    {
      label: "Try a demo",
      link: "/docs/intro",
    },
  ];

  return (
    <header className={clsx("hero hero--primary", styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          {siteConfig.title}
        </Heading>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className="flex justify-center items-center space-x-4">
          {buttons.map((button) => (
            <button className={styles.buttons}>
              <Link
                className="button button--secondary button--lg"
                to={button.link}
              >
                {button.label}
              </Link>
            </button>
          ))}
        </div>
      </div>
    </header>
  );
}

export default function Home(): JSX.Element {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`Hello from ${siteConfig.title}`}
      description="Description will go into a meta tag in <head />"
    >
      <HomepageHeader />
      <main>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
