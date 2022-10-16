import { Module } from '@nestjs/common';
import {GraphQLModule} from "@nestjs/graphql";
import {ApolloGatewayDriver, ApolloGatewayDriverConfig} from "@nestjs/apollo";
import {IntrospectAndCompose} from "@apollo/gateway";

@Module({
  imports: [
    GraphQLModule.forRoot<ApolloGatewayDriverConfig>({
      driver: ApolloGatewayDriver,
      server: {
        cors: true
      },
      gateway: {
        supergraphSdl: new IntrospectAndCompose({
          subgraphs: [
            {name: 'todos', url: 'http://federation_backend:8080/query'}
          ]
        })
      }
    })
  ],
})
export class AppModule {}
