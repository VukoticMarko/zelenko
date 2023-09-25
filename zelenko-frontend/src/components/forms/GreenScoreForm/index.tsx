import { Button, ButtonGroup, IconButton, VStack } from "@chakra-ui/react";
import { AddIcon, MinusIcon } from '@chakra-ui/icons';


const GreenScoreForm = () => {
    return (<>
        <VStack>
            <ButtonGroup size='lg' isAttached variant='outline'>
                <IconButton 
                    aria-label='Upvote' 
                    icon={<AddIcon />} 
                    onClick={() => console.log("+1")}
                />
                <IconButton 
                    aria-label='Downvote' 
                    icon={<MinusIcon/>}
                    onClick={() => console.log("-1")}
                />
            </ButtonGroup>
        </VStack>
    </>);
}

export default GreenScoreForm;