import { TabContext, TabList, TabPanel } from "@mui/lab";
import { Tab } from "@mui/material";
import { SyntheticEvent, VFC } from "react";
import { PartialDeep } from "type-fest";
import { GetClassesBySchoolQuery } from "../generated/graphql";

type Props = {
  date: string;
  school: GetClassesBySchoolQuery["school"];
  selectedStudioId: string;
  selectedRoomId: string;
  handleSelectStudioId: (studioId: string) => void;
  handleSelectRoomId: (roomId: string) => void;
  onClickClass: (clazz: ClassType) => void;
  onClickNewClass?: (schedule: PartialDeep<SchedulesType[number]>) => void;
};

type StudiosType = GetClassesBySchoolQuery["school"]["studios"];
type StudioType = StudiosType[number];
type RoomsType = StudioType["rooms"];
type RoomType = RoomsType[number];
type SchedulesType = RoomType["schedules"];
type ClassType = NonNullable<SchedulesType[number]["class"]>;
type InstructorType = ClassType["instructor"];

export const Timetable: VFC<Props> = ({
  school,
  selectedStudioId,
  selectedRoomId,
  handleSelectStudioId,
  ...props
}) => {
  const { studios } = school;
  const handleChangeStudio = (event: SyntheticEvent, newValue: string) => {
    handleSelectStudioId(newValue);
  };

  return (
    <TabContext value={selectedStudioId}>
      <TabList onChange={handleChangeStudio}>
        {studios.map((studio) => (
          <Tab key={studio.id} label={studio.name} value={studio.id} />
        ))}
      </TabList>
      {studios.map((studio) => (
        <TabPanel key={studio.id} value={studio.id} sx={{ px: 0 }}>
          <TabContext value={selectedRoomId}>
            <TabList>
              {studio.rooms.map((room) => (
                <Tab key={room.id} label={room.name} value={room.id} />
              ))}
            </TabList>
          </TabContext>
        </TabPanel>
      ))}
    </TabContext>
  );
};
