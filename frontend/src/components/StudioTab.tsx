import { VFC } from "react";
export const StudioTab: VFC = () => {
  return <></>;
};
// import { TabContext, TabList, TabPanel } from "@mui/lab";
// import { useEffect, useState, VFC, SyntheticEvent } from "react";
// import { GetClassesBySchoolQuery } from "../generated/graphql";
// import { Box, Grid, Tab, Tabs } from "@mui/material";
// import { RoomTab } from "./RoomTab";
// import { PartialDeep } from "type-fest";

// type Props = {
//   studios: StudiosType;
//   selectedStudioId?: string;
//   selectedRoomId?: string;
//   onClickClass: (clazz: ClassType) => void;
//   onClickNewClass?: (schedule: PartialDeep<SchedulesType[number]>) => void;
// };
// type StudiosType = GetClassesBySchoolQuery["school"]["studios"];
// type StudioType = StudiosType[number];
// type RoomsType = StudioType["rooms"];
// type RoomType = RoomsType[number];
// type SchedulesType = RoomType["schedules"];
// type ClassType = NonNullable<SchedulesType[number]["class"]>;
// type InstructorType = ClassType["instructor"];
// export const StudioTab: VFC<Props> = ({
//   studios,
//   selectedStudioId,
//   ...props
// }) => {
//   const [studio, setStudio] = useState<StudioType>();

//   useEffect(() => {
//     setStudio(studios[0]);
//   }, [studios]);

//   const handleChangeStudio = (event: SyntheticEvent, newValue: string) => {
//     if (studio) {
//       const newStudio = studios.find((studio) => studio.id === newValue);
//       if (newStudio) {
//         setStudio(newStudio);
//       }
//     }
//   };

//   if (!studio) return <></>;

//   return (
//     <TabContext value={studio.id}>
//       <Box>
//         <TabList onChange={handleChangeStudio}>
//           {studios.map((studio) => (
//             <Tab key={studio.id} label={studio.name} value={studio.id} />
//           ))}
//         </TabList>
//       </Box>
//       {studios.map((studio) => (
//         <TabPanel key={studio.id} value={studio.id} sx={{ px: 0 }}>
//           {/* <RoomTab rooms={studio.rooms, ...props} /> */}
//         </TabPanel>
//       ))}
//     </TabContext>
//   );
// };
