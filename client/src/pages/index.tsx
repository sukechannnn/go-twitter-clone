import React, { useState, useEffect } from "react";
import { Container, Stack } from "@chakra-ui/react";
import { Tweet, TweetForm, SideBar } from "../components/index";
import { GraphQLClient } from "graphql-request";
import { getSdk } from "../generated/graphql";

const client = new GraphQLClient("http://localhost:8080/query", {
  credentials: "include",
  mode: "cors",
});
const sdk = getSdk(client);

const Home: React.FC = () => {
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
  };

  return (
    <SideBar>
      <Container my={8}>
        <Stack spacing={8} mx={"auto"} maxW={"lg"} py={0} px={6}>
          <TweetForm
            text={tweetText}
            onChange={handleChange}
            onClick={handleButton}
          />
          <Stack spacing={3}>
            {timeline.map((tweet) => (
              <Tweet
                key={tweet.id}
                screenId={tweet.screenId}
                text={tweet.text}
              />
            ))}
          </Stack>
        </Stack>
      </Container>
    </SideBar>
  );
};

export default Home;
