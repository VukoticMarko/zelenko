import { Box, Button, Card, CardBody, CardHeader, Flex, Heading } from "@chakra-ui/react";
import { IGreenObject } from "../../common/dtos";
import { useDispatch } from "react-redux";
import { switchEdit } from "../../slices/greenObject.slice";

interface IGreenObjectDeatilsProps {
    greenObject: IGreenObject;
}

const GreenObjectDetails = ( { greenObject }: IGreenObjectDeatilsProps ) => {

    const dispatch = useDispatch();

    return (
        <Card 
            width={'85%'} 
            borderWidth={'thick'} 
            alignItems={'flex-start'} 
            padding={'10px'}
            gap={1}
        >
            <Heading size='lg'>{greenObject.LocationName}</Heading>
            <Box>
                <label>Coordinates: </label> {`[${greenObject.Location.Latitude}, ${greenObject.Location.Longitude}]`}
            </Box>
            <Box>
                <label>Address: </label> {`${greenObject.Location.Street}, ${greenObject.Location.City}, ${greenObject.Location.Country} `}
            </Box>
            <Box>
                <label>Type: </label> {greenObject.TrashType}
            </Box>
            <Box>
                <label>Shape: </label> {greenObject.Shape}
            </Box>
            <Flex
                justifyContent={'flex-end'} 
                width={"100%"}
            >
                <Button onClick={() => dispatch(switchEdit(true))}>Edit</Button>
            </Flex>
      </Card>);
}

export default GreenObjectDetails;