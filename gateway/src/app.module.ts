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
            {name: 'accounts', url: 'http://federation_accounts:8082/query'},
            {name: 'products', url: 'http://federation_products:8083/query'},
            {name: 'reviews', url: 'http://federation_reviews:8084/query'}
          ]
        })
      }
    })
  ],
})
export class AppModule {}
