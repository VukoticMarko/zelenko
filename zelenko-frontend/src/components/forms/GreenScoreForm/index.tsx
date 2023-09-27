import { ButtonGroup, Card, HStack, Heading, IconButton, VStack } from "@chakra-ui/react";
import { AddIcon, MinusIcon } from '@chakra-ui/icons';
import { IGreenObject } from "../../../common/dtos";
import GreenObjectDetails from "../../GreenObjectDetails";
import { addOneScore, subOneScore } from "../../../api/greenScore.api";
import { useDispatch } from "react-redux";
import { updateGreenScore } from "../../../slices/greenObject.slice";

interface IGreenScoreFormProps {
    greenObject: IGreenObject;
} 

const GreenScoreForm = ( { greenObject }: IGreenScoreFormProps ) => {

    const dispatch = useDispatch();
    const addOne = async () => {
        try {
            const response = await addOneScore(greenObject);
            dispatch(updateGreenScore(response));
        } catch (err) {
            console.log(err);
        }
    }

    const subOne = async () => {
        try {
            const response = await subOneScore(greenObject);
            dispatch(updateGreenScore(response));
        } catch (err) {
            console.log(err);
        }
    }

    return (
    <HStack marginTop={'4'}>
        <GreenObjectDetails greenObject={greenObject}/>
        <Card borderWidth={'thick'}>
            <VStack>
                <Heading as='h1' size={'4xl'}>{greenObject.GreenScore.Verification}</Heading>
                <label>{greenObject.GreenScore.TrashRank}</label>
                <ButtonGroup size='lg' isAttached variant='outline'>
                    <IconButton 
                        aria-label='Upvote' 
                        icon={<AddIcon />} 
                        onClick={addOne}
                    />
                    <IconButton 
                        aria-label='Downvote' 
                        icon={<MinusIcon/>}
                        onClick={subOne}
                    />
                </ButtonGroup>
            </VStack>
        </Card>
    </HStack>);
}

export default GreenScoreForm;