import { Button, Stack, Textarea } from "@chakra-ui/react";
import React from "react";

type Props = {
  text: string;
  onChange: (e) => void;
  onClick: (e) => void;
};

export const TweetForm: React.FC<Props> = ({ text, onChange, onClick }) => {
  return (
    <Stack
      border="1px"
      borderRadius="md"
      borderColor="gray.200"
      p={4}
      bg="white"
    >
      <Textarea
        value={text}
        onChange={onChange}
        placeholder="いまどうしてる？"
        size="sm"
      />
      <Button
        bg={"blue.400"}
        color={"white"}
        _hover={{
          bg: "blue.500",
        }}
        onClick={onClick}
      >
        Tweet
      </Button>
    </Stack>
  );
};
