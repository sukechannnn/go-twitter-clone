import { Avatar, Box, Button, Heading, HStack, Stack } from "@chakra-ui/react";
import React from "react";

type Props = {
  id: string;
  screenId: string;
  screenName: string;
  following: boolean;
  onClick: (e) => void;
};

export const User: React.FC<Props> = ({
  id,
  screenId,
  screenName,
  following,
  onClick,
}) => {
  return (
    <HStack
      border="1px"
      borderRadius="md"
      borderColor="gray.200"
      py={4}
      px={8}
      bg="white"
      justifyContent={"space-between"}
    >
      <HStack>
        <Avatar mr={4} />
        <Stack>
          <Heading size="sm">@{screenId}</Heading>
          <Box>{screenName}</Box>
        </Stack>
      </HStack>
      <Button
        id={id}
        bg={"green.300"}
        color={"white"}
        _disabled={{
          bg: "green.100",
        }}
        _hover={{
          bg: "green.400",
          _disabled: {
            bg: "green.100",
            cursor: "default",
          },
        }}
        disabled={following}
        onClick={onClick}
      >
        {following ? "following" : "follow"}
      </Button>
    </HStack>
  );
};
