import * as React from "react"
import {
  ChakraProvider,
  Box,
  Text,
  VStack,
  Code,
  Grid,
  theme,
  Center,
} from "@chakra-ui/react"
import GreenObjectsPage from "./pages/GreenObjectsPage"
import { Provider } from "react-redux"
import { store } from "./slices/store";

export const App = () => (
  <ChakraProvider theme={theme}>
    <Provider store={store}>
    <Box textAlign="center" fontSize="xl">
      <Grid minH="100vh" p={3}>
        <VStack spacing={8}>
          <Text>
            Edit <Code fontSize="xl">src/App.tsx</Code> and save to reload.
          </Text>
          <Center>
            <GreenObjectsPage/>
          </Center>
        </VStack>
      </Grid>
    </Box>
    </Provider>
  </ChakraProvider>
)
