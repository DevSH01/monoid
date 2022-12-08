import {
  gql, useQuery,
} from '@apollo/client';
import { CircleStackIcon, FolderIcon } from '@heroicons/react/24/outline';
import AlertRegion from 'components/AlertRegion';
import Badge from 'components/Badge';
import Button from 'components/Button';
import { CardDivider } from 'components/Card';
import EmptyState from 'components/Empty';
import { MonoidA } from 'components/MonoidLink';
import Spinner from 'components/Spinner';
import Table from 'components/Table';
import Text from 'components/Text';
import { DataSource, Property, SiloDefinition } from 'lib/models';
import React, { forwardRef, useImperativeHandle } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { dedup } from 'utils/utils';
import CategoryCombobox from './CategoryCombobox';
import IdentifierSelect from './IdentifierSelect';

const SILO_DATA_SOURCES = gql`
  query SiloDataSources($id: ID!) {
    siloDefinition(id: $id) {
      id
      dataSources {
        id
        name
        group
        properties {
          id
          name
          categories {
            id
            name
          }
          userPrimaryKey {
            id
            name
          }
        }
      }
    }
  }
`;

const DataSourcesTable = forwardRef((props: { siloId: string, type: 'card' | 'plain' }, ref: React.ForwardedRef<{
  refetch: () => void
} | undefined>) => {
  const { siloId, type } = props;
  const { id } = useParams<{ id: string }>();

  const {
    data, loading, error, refetch,
  } = useQuery<{ siloDefinition: SiloDefinition }>(SILO_DATA_SOURCES, {
    variables: {
      id: siloId,
    },
  });
  const navigate = useNavigate();

  useImperativeHandle(ref, () => ({
    refetch() {
      if (refetch) refetch();
    },
  }));

  if (loading) {
    return (
      <div className="md:px-6 md:-mt-6 px-4 -mt-5 md:pb-6 pb-4">
        <Spinner />
      </div>
    );
  }

  if (error) {
    return (
      <div className="md:px-6 md:-mt-6 px-4 -mt-5 md:pb-6 pb-4">
        <AlertRegion alertTitle="Error">
          {error.message}
        </AlertRegion>
      </div>
    );
  }

  const empty = !data?.siloDefinition?.dataSources?.length;

  if (!empty) {
    return (
      <Table
        type={type}
        tableCols={[
          {
            header: 'Name',
            key: 'name',
          },
          {
            header: 'Tracked Properties',
            key: 'properties',
          },
          {
            header: (
              <div className="flex-flex-col">
                <div>Purpose of Processing</div>
                <Text em="light">
                  Requires
                  <MonoidA href="https://monoid.co" className="ml-1 underline" target="_blank">a license</MonoidA>
                </Text>
              </div>
            ),
            key: 'purposes',
          },
        ]}
        tableRows={data?.siloDefinition?.dataSources?.map((ds: DataSource) => ({
          key: ds.id!,
          columns: [
            {
              key: 'name',
              content: (
                <div className="flex flex-col items-start">
                  <div>
                    <div>{ds.name}</div>
                    <Text size="xs" em="light" className="flex items-center">
                      <FolderIcon className="w-3 h-3 mr-1" />
                      {ds.group}
                    </Text>
                  </div>
                </div>
              ),
            },
            {
              key: 'properties',
              content: (
                <div className="space-x-2">
                  {
                    dedup(
                      // Get all the categories that are listed under the properties for
                      // the data source.
                      ds.properties?.flatMap((p) => p.categories || []) || [],
                      (c) => c.id!,
                    ).map((c) => (
                      <Badge key={c.id} color="blue">
                        {c.name}
                      </Badge>
                    ))
                  }
                </div>
              ),
            },
            {
              key: 'purposes',
              content: null,
            },
          ],

          nestedComponent: (
            <tr>
              <td colSpan={4} className="p-0">
                <div>
                  <Table
                    tableCols={[{
                      header: 'Property Name',
                      key: 'name',
                    }, {
                      header: 'Category',
                      key: 'cat',
                    }, {
                      header: 'User Identifier',
                      key: 'user_identifier',
                    }]}
                    tableRows={
                      ds.properties?.map(
                        (p: Property) => ({
                          key: p.id!,
                          columns: [
                            {
                              key: 'name',
                              content: (
                                <div className="flex flex-col items-start">
                                  {p.name}
                                </div>
                              ),
                            },
                            {
                              key: 'cat',
                              content: (
                                <CategoryCombobox
                                  value={p.categories?.map((c) => c.id!) || []}
                                  propertyId={p.id!}
                                />
                              ),
                            },
                            {
                              key: 'identifier',
                              content: (
                                <IdentifierSelect
                                  value={p.userPrimaryKey?.id}
                                  workspaceId={id!}
                                  propertyId={p.id!}
                                />
                              ),
                            },
                          ],
                        }),
                      )
                    }
                    type="plain"
                    insetClass="pl-12"
                    className="border-y-2 border-gray-300"
                  />
                </div>
              </td>
            </tr>
          ),
        }))}
        innerClassName="border-t border-gray-300"
        nested
      />
    );
  }

  return (
    <div className="md:px-6 md:-mt-6 px-4 -mt-5 md:pb-6 pb-4">
      <CardDivider />
      <EmptyState
        icon={CircleStackIcon}
        title="No Data Sources"
        subtitle="You can find data sources by running a scan and applying alerts."
        action={(
          <Button onClick={() => navigate('../alerts')}>
            View Alerts
          </Button>
        )}
        className="py-8"
      />
    </div>
  );
});

export default DataSourcesTable;
