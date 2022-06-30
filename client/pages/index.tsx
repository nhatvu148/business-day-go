import { Button } from "@mui/material";
import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";

const Home: NextPage = () => {
  return (
    <div>
      <Head>
        <title>Business Day</title>
        <meta name="description" content="Business Day" />
        <meta name="viewport" content="initial-scale=1, width=device-width" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div>Testing fonts</div>
      <Button variant="contained">Hello World</Button>
    </div>
  );
};

export default Home;
