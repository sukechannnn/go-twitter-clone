import React from "react";
import { Container, Heading, Stack } from "@chakra-ui/react";
import { Tweet } from "../components/tweet";

const Home = () => {
  return (
    <Container my={8}>
      <Heading>Twitter clone!</Heading>
      <Stack spacing={4} mx={"auto"} maxW={"lg"} py={12} px={6}>
        <Tweet screenId="Hoge" text="Hello world!" />
        <Tweet screenId="Hoge" text="Hello world!" />
        <Tweet screenId="Hoge" text="Hello world!" />
      </Stack>
    </Container>
  );
};

export default Home;
