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

export type GetTimelineQueryVariables = Exact<{ [key: string]: never; }>;


export type GetTimelineQuery = { __typename?: 'Query', timeline: Array<{ __typename?: 'PostInfo', id: string, userId: string, text: string, screenId: string, createdAt: any }> };


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

export type SdkFunctionWrapper = <T>(action: (requestHeaders?:Record<string, string>) => Promise<T>, operationName: string) => Promise<T>;


const defaultWrapper: SdkFunctionWrapper = (action, _operationName) => action();

export function getSdk(client: GraphQLClient, withWrapper: SdkFunctionWrapper = defaultWrapper) {
  return {
    getTimeline(variables?: GetTimelineQueryVariables, requestHeaders?: Dom.RequestInit["headers"]): Promise<GetTimelineQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetTimelineQuery>(GetTimelineDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'getTimeline');
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;