import React, { useState, useEffect } from "react";
import { Container, Heading, Stack } from "@chakra-ui/react";
import { Tweet, TweetForm } from "../components/index";
import { GraphQLClient } from "graphql-request";
import { getSdk } from "../generated/graphql";

const client = new GraphQLClient("http://localhost:8080/query", {
  credentials: "include",
  mode: "cors",
});
const sdk = getSdk(client);

const Home = () => {
  const [tweetText, setTweetText] = useState("");
  const [timeline, setTimeline] = useState([]);
  useEffect(() => {
    sdk.getTimeline().then((res) => setTimeline(res.timeline));
  }, []);

  const handleChange = (event) => {
    setTweetText(event.target.value);
  };

  const handleButton = (event) => {
    event.preventDefault();
    sdk
      .tweet({ text: tweetText })
      .then(() => sdk.getTimeline().then((res) => setTimeline(res.timeline)));
    setTweetText("");
    console.log(timeline);
  };

  return (
    <Container my={8}>
      <Stack spacing={8} mx={"auto"} maxW={"lg"} py={6} px={6}>
        <Heading>Twitter clone!</Heading>
        <TweetForm
          text={tweetText}
          onChange={handleChange}
          onClick={handleButton}
        />
        <Stack spacing={4}>
          {timeline.map((tweet) => (
            <Tweet key={tweet.id} screenId={tweet.screenId} text={tweet.text} />
          ))}
        </Stack>
      </Stack>
    </Container>
  );
};

export default Home;
