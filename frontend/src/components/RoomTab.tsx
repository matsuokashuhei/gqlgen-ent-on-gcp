import { VFC } from "react";
export const RoomTab: VFC = () => {
  return <></>;
};
// import { TabContext, TabList, TabPanel } from "@mui/lab";
// import { useEffect, useState, VFC, SyntheticEvent } from "react";
// import { GetClassesBySchoolQuery } from "../generated/graphql";
// import { Box, Grid, Tab, Tabs } from "@mui/material";

// type StudiosType = GetClassesBySchoolQuery["school"]["studios"];
// type StudioType = StudiosType[number];
// type RoomsType = StudioType["rooms"];
// type RoomType = RoomsType[number];

// type Props = {
//   rooms: RoomsType;
//   selectedRoomId?: string;
//   onClickClass: (clazz: ClassType) => void;
//   onClickNewClass?: (schedule: PartialDeep<SchedulesType[number]>) => void;
// };

// export const RoomTab: VFC<Props> = ({ rooms }) => {
//   const [room, setRoom] = useState<RoomType>();

//   useEffect(() => {
//     setRoom(rooms[0]);
//   }, [rooms]);

//   const handleChangeRoom = (event: SyntheticEvent, newValue: string) => {
//     if (room) {
//       const newRoom = rooms.find((room) => room.id === newValue);
//       if (newRoom) {
//         setRoom(newRoom);
//       }
//     }
//   };

//   if (!room) return <></>;

//   return (
//     <TabContext value={room.id}>
//       <Box>
//         <TabList onChange={handleChangeRoom}>
//           {rooms.map((room) => (
//             <Tab key={room.id} label={room.name} value={room.id} />
//           ))}
//         </TabList>
//       </Box>
//       {rooms.map((room) => (
//         <TabPanel key={room.id} value={room.id}>
//           {room.name}
//         </TabPanel>
//       ))}
//     </TabContext>
//   );
// };
