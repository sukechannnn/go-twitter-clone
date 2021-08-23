import { GraphQLClient } from 'graphql-request';
import * as Dom from 'graphql-request/dist/types.dom';
import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
};

export type FollowUser = {
  __typename?: 'FollowUser';
  id: Scalars['ID'];
  userId: Scalars['String'];
  followId: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createUser: User;
  followUser: FollowUser;
  createPost: Post;
};


export type MutationCreateUserArgs = {
  input: NewUser;
};


export type MutationFollowUserArgs = {
  input: NewFollowUser;
};


export type MutationCreatePostArgs = {
  input: NewPost;
};

export type NewFollowUser = {
  followId: Scalars['String'];
};

export type NewPost = {
  text: Scalars['String'];
};

export type NewUser = {
  email: Scalars['String'];
  password: Scalars['String'];
  screenId: Scalars['String'];
  screenName: Scalars['String'];
};

export type Post = {
  __typename?: 'Post';
  id: Scalars['ID'];
  userId: Scalars['String'];
  text: Scalars['String'];
  createdAt: Scalars['Time'];
};

export type PostInfo = {
  __typename?: 'PostInfo';
  id: Scalars['ID'];
  userId: Scalars['String'];
  text: Scalars['String'];
  screenId: Scalars['String'];
  createdAt: Scalars['Time'];
};

export type Query = {
  __typename?: 'Query';
  allUsers: Array<UserInfo>;
  user: User;
  followUsers: Array<User>;
  followers: Array<User>;
  timeline: Array<PostInfo>;
};


export type QueryUserArgs = {
  id: Scalars['ID'];
};


export type User = {
  __typename?: 'User';
  id: Scalars['ID'];
  email: Scalars['String'];
  screenId: Scalars['String'];
  screenName: Scalars['String'];
  createdAt: Scalars['Time'];
};

export type UserInfo = {
  __typename?: 'UserInfo';
  id: Scalars['ID'];
  email: Scalars['String'];
  screenId: Scalars['String'];
  screenName: Scalars['String'];
  following: Scalars['Boolean'];
  createdAt: Scalars['Time'];
};

export type FollowUserMutationVariables = Exact<{
  userId: Scalars['String'];
}>;


export type FollowUserMutation = { __typename?: 'Mutation', followUser: { __typename?: 'FollowUser', id: string, followId: string, userId: string } };

export type GetTimelineQueryVariables = Exact<{ [key: string]: never; }>;


export type GetTimelineQuery = { __typename?: 'Query', timeline: Array<{ __typename?: 'PostInfo', id: string, userId: string, text: string, screenId: string, createdAt: any }> };

export type GetUsersQueryVariables = Exact<{ [key: string]: never; }>;


export type GetUsersQuery = { __typename?: 'Query', allUsers: Array<{ __typename?: 'UserInfo', id: string, screenId: string, screenName: string, following: boolean }> };

export type TweetMutationVariables = Exact<{
  text: Scalars['String'];
}>;


export type TweetMutation = { __typename?: 'Mutation', createPost: { __typename?: 'Post', id: string, text: string, createdAt: any, userId: string } };


export const FollowUserDocument = gql`
    mutation followUser($userId: String!) {
  followUser(input: {followId: $userId}) {
    id
    followId
    userId
  }
}
    `;
export const GetTimelineDocument = gql`
    query getTimeline {
  timeline {
    id
    userId
    text
    screenId
    createdAt
  }
}
    `;
export const GetUsersDocument = gql`
    query getUsers {
  allUsers {
    id
    screenId
    screenName
    following
  }
}
    `;
export const TweetDocument = gql`
    mutation tweet($text: String!) {
  createPost(input: {text: $text}) {
    id
    text
    createdAt
    userId
  }
}
    `;

export type SdkFunctionWrapper = <T>(action: (requestHeaders?:Record<string, string>) => Promise<T>, operationName: string) => Promise<T>;


const defaultWrapper: SdkFunctionWrapper = (action, _operationName) => action();

export function getSdk(client: GraphQLClient, withWrapper: SdkFunctionWrapper = defaultWrapper) {
  return {
    followUser(variables: FollowUserMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<FollowUserMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<FollowUserMutation>(FollowUserDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'followUser');
    },
    getTimeline(variables?: GetTimelineQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<GetTimelineQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetTimelineQuery>(GetTimelineDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'getTimeline');
    },
    getUsers(variables?: GetUsersQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<GetUsersQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetUsersQuery>(GetUsersDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'getUsers');
    },
    tweet(variables: TweetMutationVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<TweetMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<TweetMutation>(TweetDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'tweet');
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;