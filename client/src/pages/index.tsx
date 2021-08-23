import React, { useState, useEffect } from "react";
import { Container, Heading, Stack } from "@chakra-ui/react";
import { Tweet } from "../components/tweet";
import { GraphQLClient } from "graphql-request";
import { getSdk } from "../generated/graphql";

const client = new GraphQLClient("http://localhost:8080/query", {
  credentials: "include",
  mode: "cors",
});
const sdk = getSdk(client);

const Home = () => {
  const [timeline, setTimeline] = useState([]);
  useEffect(() => {
    sdk.getTimeline().then((res) => setTimeline(res.timeline));
  }, []);

  return (
    <Container my={8}>
      <Heading>Twitter clone!</Heading>
      <Stack spacing={4} mx={"auto"} maxW={"lg"} py={12} px={6}>
        {timeline.map((tweet) => (
          <Tweet key={tweet.id} screenId={tweet.screenId} text={tweet.text} />
        ))}
      </Stack>
    </Container>
  );
};

export default Home;
