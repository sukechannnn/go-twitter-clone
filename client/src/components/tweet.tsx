import { Avatar, Box, Heading, HStack, Stack } from "@chakra-ui/react";
import React from "react";

type Props = {
  screenId: string;
  text: string;
};

export const Tweet: React.FC<Props> = ({ screenId, text }) => {
  return (
    <HStack border="1px" borderRadius="md" borderColor="gray.200" p={4}>
      <Avatar mr={4} />
      <Stack>
        {" "}
        <Heading size="sm">{screenId}</Heading>
        <Box>{text}</Box>
      </Stack>
    </HStack>
  );
};
