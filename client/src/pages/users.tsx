import { chakra, Container, Heading, Stack, useId } from "@chakra-ui/react";
import React, { useEffect, useState } from "react";
import { SideBar, User } from "src/components";
import { GraphQLClient } from "graphql-request";
import { getSdk } from "../generated/graphql";

const client = new GraphQLClient("http://localhost:8080/query", {
  credentials: "include",
  mode: "cors",
});
const sdk = getSdk(client);

const Users: React.FC = () => {
  const [users, setUsers] = useState([]);
  const [followId, setFollowId] = useState("");
  useEffect(() => {
    sdk.getUsers().then((res) => setUsers(res.allUsers));
  }, []);

  const handleButton = (event) => {
    event.preventDefault();
    sdk.followUser({ userId: event.target.id }).then(() => {
      sdk.getUsers().then((res) => setUsers(res.allUsers));
    });
  };

  return (
    <SideBar>
      <Container my={8}>
        <Stack spacing={8} mx={"auto"} maxW={"lg"} py={0} px={6}>
          <Heading size={"lg"}>ユーザー一覧</Heading>
          <Stack spacing={3}>
            {users.map((user) => (
              <User
                key={user.id}
                id={user.id}
                screenId={user.screenId}
                screenName={user.screenName}
                following={user.following}
                onClick={handleButton}
              />
            ))}
          </Stack>
        </Stack>
      </Container>
    </SideBar>
  );
};

export default Users;
